package ports

type Hasher interface {
  Hash(string) (string, error)
  CompareHashAndPasswd(string, string) bool
}
