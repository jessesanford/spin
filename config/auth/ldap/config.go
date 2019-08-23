package ldap

type LdapConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (l *LdapConfig) IsValid() bool {
  //noop
  return true
}
