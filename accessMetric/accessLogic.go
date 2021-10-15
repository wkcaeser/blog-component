package accessMetric

import (
	db "blog-component/db"
)

type AccessLog struct {
	Id         int64  `json:"id"`
	UserUuid   string `json:"user_uuid"`
	Uri        string `json:"uri"`
	Site       string `json:"site"`
	ReqUri     string `json:"req_uri"`
	Ip         string `json:"ip"`
	RequestLog string `json:"request_log"`
}

type PvMetric struct {
	PeopleTotal int64 `json:"peopleTotal"`
	PeoplePv    int64 `json:"peoplePv"`
}

type AccessLogImpl struct {
}

type IAccessLog interface {
	GetUriMetric(url string) PvMetric
	PutUriMetric(metric *AccessLog)
}

func (*AccessLogImpl) GetUriMetric(uri string, uuid string) PvMetric {

	var pvMetric PvMetric

	stmtPv := db.Stmt("SELECT count(*) as PeoplePv from access_metric WHERE uri = ?")

	defer stmtPv.Close()
	stmtPv.QueryRow(uri).Scan(&pvMetric.PeoplePv)

	stmtTotal := db.Stmt("SELECT count(distinct user_uuid) as PeopleTotal cnt from access_metric")

	defer stmtTotal.Close()
	stmtTotal.QueryRow(uri).Scan(&pvMetric.PeopleTotal)

	return pvMetric
}

func (*AccessLogImpl) PutUrlMetric(accessLog *AccessLog) {
	stmt := db.Stmt("insert into access_log (user_uuid, uri, site, req_uri, ip, request_log) values (?, ?, ?, ?, ?, ?)")
	stmt.Exec(accessLog.UserUuid, accessLog.Uri, accessLog.Site, accessLog.ReqUri, accessLog.Ip, accessLog.RequestLog)
}
