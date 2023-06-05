package webhook

import (
	resty "github.com/go-resty/resty/v2"
	"github.com/gotomicro/ego/core/elog"

	"github.com/kl7sn/apian/pkg/dto"
	"github.com/kl7sn/apian/pkg/utils"
)

func Webhook(fw dto.Webhook) {
	client := resty.New()
	r := client.R().SetHeader("Content-Type", "application/json")
	for k, v := range fw.Headers {
		r.SetHeader(k, v)
	}
	if _, ok := fw.Body["mode"]; ok {
		ip, err := utils.GetOutBoundIP()
		if err != nil {
			elog.Error("forward", elog.FieldErr(err), elog.String("msg", "get outbound ip error"))
			return
		}
		fw.Body["addr"] = ip + ":9003"
		elog.Info("apian", elog.String("msg", "forward success"), elog.String("ip", ip), elog.Any("body", fw.Body))
	}
	r.SetBody(fw.Body)
	r.Post(fw.Url)
}