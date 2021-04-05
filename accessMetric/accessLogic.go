package accessMetric

import db "blog-component/db"

type AccessMetric struct {
	Id            int64  `json:"id"`
	Url           string `json:"url"`
	BrowseCnt     int64  `json:"browse_cnt"`
	BrowseUserCnt int64  `json:"browse_user_cnt"`
	Comment       string `json:"comment"`
}

type AccessLogicImpl struct {
}

type AccessLogic interface {
	GetUrlMetric(url string) *AccessMetric
	PutUrlMetric(metric *AccessMetric)
}

func (*AccessLogicImpl) GetUrlMetric(url string) *AccessMetric {

	metric := &AccessMetric{}

	stmt := db.Stmt("SELECT id, url, browse_cnt, browse_user_cnt, comment from access_metric WHERE url = ?")

	defer stmt.Close()

	row := stmt.QueryRow(url)
	if row != nil {
		row.Scan(&metric.Id, &metric.Url, &metric.BrowseCnt, &metric.BrowseUserCnt, &metric.Comment)
	}

	return metric
}

func (impl AccessLogicImpl) PutUrlMetric(metric *AccessMetric) {

	if metric.Url == "" {
		return
	}

	dbSavedMetric := impl.GetUrlMetric(metric.Url)

	if dbSavedMetric.Id == 0 {
		stmt := db.Stmt("insert into access_metric(url, browse_cnt, browse_user_cnt, comment) values (?, ?, ?, ?)")
		stmt.Exec(metric.Url, metric.BrowseCnt, metric.BrowseUserCnt, metric.Comment)
	} else {
		stmt := db.Stmt("update access_metric set browse_cnt = browse_cnt + ?, browse_user_cnt = browse_user_cnt + ? where id = ?")
		stmt.Exec(metric.BrowseCnt, metric.BrowseUserCnt, dbSavedMetric.Id)
	}
}
