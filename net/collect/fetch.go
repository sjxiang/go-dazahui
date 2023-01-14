
// 采集

package collect


type Fetcher interface {
	Get(url string) ([]byte, error)
}

