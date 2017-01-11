package config

type Config struct {
	HTTP       *HTTPConfig       `toml:"http"`
	Logs       *LogsConfig       `toml:"logs"`
	Credential *CredentialConfig `toml:"credential"`
	Storage    *StorageConfig    `toml:"storage"`
	Init       *InitConfig       `toml:"init"`
}

type HTTPConfig struct {
	HTTPAddr string `toml:"http_addr"`
	Scheme   string `toml:"scheme"`
}

type LogsConfig struct {
	AccessLog bool   `toml:"access_log"`
	Path      string `toml:"path"`
	Level     string `toml:"level"`
}

type CredentialConfig struct {
	AccessKey string `toml:"accesskey"`
	SecretKey string `toml:"secketkey"`
}

type StorageConfig struct {
	Type  string       `toml:"type"`
	Swift *SwiftConfig `toml:"swift"`
	S3    *S3Config    `toml:"s3"`
	OSS   *OSSConfig   `toml:"oss"`
	Posix *PosixConfig `toml:"posix"`
	Kdfs  *KdfsConfig  `toml:"kdfs"`
}

type SwiftConfig struct {
	Addr       string `toml:"addr"`
	AuthURL    string `toml:"authurl"`
	TenantName string `toml:"tenantname"`
	UserName   string `toml:"username"`
	ProjName   string `toml:"projname"`
	PassWord   string `toml:"password"`
}

type KdfsConfig struct {
	Addr    string `toml:"addr"`
	Account string `toml:"account"`
}

type S3Config struct {
	Addr      string `toml:"addr"`
	AccessKey string `toml:"accesskey"`
	SecretKey string `toml:"secretkey"`
}

type OSSConfig struct {
	Addr      string `toml:"addr"`
	AccessKey string `toml:"accesskey"`
	SecretKey string `toml:"secretkey"`
}

type PosixConfig struct {
	Addr string `toml:"addr"`
}

type InitConfig struct {
	Buckets   []string `toml:"buckets"`
	UploadDir string   `toml:"uploaddir"`
}
