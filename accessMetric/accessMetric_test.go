package accessMetric

import (
	"fmt"
	"testing"
)

func TestAccessLogicImpl_GetUrlMetric(t *testing.T) {

	var metricOp AccessLogicImpl
	metric := metricOp.GetUrlMetric("test")

	fmt.Println(metric)

}

func TestAccessLogicImpl_PutUrlMetric(t *testing.T) {

	var metricOp AccessLogicImpl
	metric := metricOp.GetUrlMetric("test")

	metric.BrowseCnt = 1
	metricOp.PutUrlMetric(metric)

	fmt.Println(metricOp.GetUrlMetric("test"))
}
