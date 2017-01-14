package middleware

import (
	"bytes"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/flyaways/storage/config"
	errors "github.com/flyaways/storage/errors"
	"github.com/flyaways/storage/result"
	"github.com/flyaways/storage/util"
	"github.com/flyaways/storage/util/log"
	"github.com/gin-gonic/gin"
)

type SignAuth struct {
	context           *gin.Context
	credential        *config.CredentialConfig
	formattedTime     string
	canonicalHeaders  string
	canonicalResource string
	stringToSign      string
	signature         string
}

func Authority(credential *config.CredentialConfig) gin.HandlerFunc {
	return (&SignAuth{credential: credential}).Auth
}

const (
	authHeaderPrefix    = "KSS"
	amzHeaderPrefix     = "x-amz-"
	kssHeaderPrefix     = "x-kss-"
	userMetaPrefix      = "x-kss-meta-"
	paramSignature      = "Signature"
	paramExpires        = "Expires"
	paramKSSAccessKeyId = "KSSAccessKeyId"
	paramAWSAccessKeyId = "AWSAccessKeyId"
	timeFormat          = "Mon, 02 Jan 2006 15:04:05 GMT"
)

var signQuerys = map[string]bool{
	"acl":                          true,
	"lifecycle":                    true,
	"location":                     true,
	"logging":                      true,
	"notification":                 true,
	"policy":                       true,
	"requestPayment":               true,
	"torrent":                      true,
	"uploadId":                     true,
	"uploads":                      true,
	"versionId":                    true,
	"versioning":                   true,
	"versions":                     true,
	"website":                      true,
	"delete":                       true,
	"thumbnail":                    true,
	"cors":                         true,
	"pfop":                         true,
	"querypfop":                    true,
	"adp":                          true,
	"queryadp":                     true,
	"partNumber":                   true,
	"response-content-type":        true,
	"response-content-language":    true,
	"response-expires":             true,
	"response-cache-control":       true,
	"response-content-disposition": true,
	"response-content-encoding":    true,
}

func (s *SignAuth) Auth(ctx *gin.Context) {
	s.context = ctx
	accessKey, signString := s.getOldSign()
	if signString == "" || accessKey != s.credential.AccessKey {
		ctx.AbortWithStatus(http.StatusForbidden)
		result.NewResult(ctx).Error(errors.SignatureDoesNotMatch)
		log.Error("[credential.AccessKey,accessKey:%s,%s]", s.credential.AccessKey, accessKey)
		return
	}

	if err := s.sign(); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		result.NewResult(ctx).Error(err)
		log.Error("[sign:%s]", err.Error())
		return
	}
	if s.signature != signString {
		ctx.AbortWithStatus(http.StatusForbidden)
		result.NewResult(ctx).Error(errors.SignatureDoesNotMatch)
		log.Error("[signature,signString:%s,%s]", s.signature, signString)
		return
	}
}

func (s *SignAuth) sign() error {
	if err := s.buildTime(); err != nil {
		return err
	}
	s.buildCanonicalHeaders()
	s.buildCanonicalResource()
	if err := s.buildStringToSign(); err != nil {
		return errors.SignatureDoesNotMatch
	}
	s.buildSignature()
	return nil
}

func (s *SignAuth) buildTime() error {

	headerDate := s.context.Request.Header.Get("Date")
	if headerDate != "" {
		s.formattedTime = headerDate
		return nil
	}
	if expireTime, ok := s.getValues(paramExpires); ok {
		if i, err := strconv.ParseInt(expireTime, 10, 0); err == nil {
			if time.Now().Unix() > i {
				return errors.URLExpired
			}
			s.formattedTime = expireTime
			return nil
		}
	}
	return nil
}

func (s *SignAuth) buildCanonicalHeaders() {
	var headers []string
	for k := range s.context.Request.Header {
		if strings.HasPrefix(strings.ToLower(http.CanonicalHeaderKey(k)), kssHeaderPrefix) {
			headers = append(headers, k)
		}
	}
	for k := range s.context.Request.Header {
		if strings.HasPrefix(strings.ToLower(http.CanonicalHeaderKey(k)), amzHeaderPrefix) {
			headers = append(headers, k)
		}
	}
	sort.Strings(headers)
	headerValues := make([]string, len(headers))
	for i, k := range headers {
		headerValues[i] = strings.ToLower(http.CanonicalHeaderKey(k)) + ":" +
			strings.Join(s.context.Request.Header[http.CanonicalHeaderKey(k)], ",")
	}
	s.canonicalHeaders = strings.Join(headerValues, "\n")
}

func (s *SignAuth) buildCanonicalResource() {
	s.context.Request.URL.RawQuery = strings.Replace(s.context.Request.URL.Query().Encode(), "+", "%20", -1)
	uri := s.context.Request.URL.Path
	uris := strings.Split(uri, "/")
	if len(uris) == 1 && uris[0] != "" {
		uri += "/"
	}
	if uri == "" {
		uri = "/"
	}
	uri = EscapePath(uri, false)
	var querys []string
	for k := range s.context.Request.URL.Query() {
		if _, ok := signQuerys[k]; ok {
			querys = append(querys, k)
		}
	}
	sort.Strings(querys)
	queryValues := make([]string, len(querys))
	for i, k := range querys {
		v := s.context.Request.URL.Query()[k]
		vString := strings.Join(v, ",")
		if vString != "" {
			queryValues[i] = k + "=" + vString
		} else {
			queryValues[i] = k
		}
	}
	queryString := strings.Join(queryValues, "&")
	if queryString == "" {
		s.canonicalResource = uri
	} else {
		s.canonicalResource = uri + "?" + queryString
	}
}

func (s *SignAuth) buildStringToSign() error {
	if s.context.Request.Method == "POST" {
		if policy, ok := s.getValues("Policy", "policy"); ok {
			s.stringToSign = policy
			return nil
		}
	}

	md5list := s.context.Request.Header["Content-Md5"]
	md5 := ""
	if len(md5list) > 0 {
		md5 = s.context.Request.Header["Content-Md5"][0]
	}

	typelist := s.context.Request.Header["Content-Type"]
	contenttype := ""
	if len(typelist) > 0 {
		contenttype = s.context.Request.Header["Content-Type"][0]
	}

	signItems := []string{s.context.Request.Method, md5, contenttype}
	signItems = append(signItems, s.formattedTime)
	if s.canonicalHeaders != "" {
		signItems = append(signItems, s.canonicalHeaders)
	}
	signItems = append(signItems, s.canonicalResource)

	s.stringToSign = strings.Join(signItems, "\n")
	return nil
}

func (s *SignAuth) buildSignature() {
	secret := s.credential.SecretKey
	signature := string(util.Base64Encode(util.MakeHmac([]byte(secret), []byte(s.stringToSign))))
	s.signature = signature
}

func (s *SignAuth) getValues(keys ...string) (string, bool) {
	for _, key := range keys {

		if value, ok := s.context.GetQuery(key); ok {
			return value, true
		}

		if value, ok := s.context.GetPostForm(key); ok {
			return value, true
		}
	}
	return "", false
}

func (s *SignAuth) getOldSign() (string, string) {
	if qs, ok := s.getValues("Signature", "signature"); ok {
		if kssKey, ok := s.getValues(paramKSSAccessKeyId, paramAWSAccessKeyId); ok {
			return kssKey, qs
		}
	}
	hs := s.context.Request.Header.Get("Authorization")
	if hs != "" {
		hs := hs[4:]
		ks := strings.Split(hs, ":")
		return ks[0], ks[1]
	}
	return "", ""
}

var noEscape [256]bool

func init() {
	for i := 0; i < len(noEscape); i++ {
		noEscape[i] = (i >= 'A' && i <= 'Z') ||
			(i >= 'a' && i <= 'z') ||
			(i >= '0' && i <= '9') ||
			i == '-' ||
			i == '.' ||
			i == '_' ||
			i == '~'
	}
}

func EscapePath(path string, encodeSep bool) string {
	var buf bytes.Buffer
	for i := 0; i < len(path); i++ {
		c := path[i]
		if noEscape[c] || (c == '/' && !encodeSep) {
			buf.WriteByte(c)
		} else {
			buf.WriteByte('%')
			buf.WriteString(strings.ToUpper(strconv.FormatUint(uint64(c), 16)))
		}
	}
	return buf.String()
}
