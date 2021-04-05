let host = window.location.host;
let uri = window.location.pathname;
let apiHost = '127.0.0.1'

let pageMetricHttpRequest = new XMLHttpRequest();//第一步：建立所需的对象
pageMetricHttpRequest.open('GET', apiHost + '/browse/page?url=' + uri, true);//第二步：打开连接  将请求参数写在url中  ps:"./Ptest.php?name=test&nameone=testone"
pageMetricHttpRequest.send();//第三步：发送请求  将请求参数写在URL中

let siteMetricHttpRequest = new XMLHttpRequest();//第一步：建立所需的对象
siteMetricHttpRequest.open('GET', apiHost + '/browse/page?url=' + uri, true);//第二步：打开连接  将请求参数写在url中  ps:"./Ptest.php?name=test&nameone=testone"
siteMetricHttpRequest.send();//第三步：发送请求  将请求参数写在URL中


let metricRequest = new XMLHttpRequest();//第一步：建立所需的对象
metricRequest.open('GET', apiHost + '/pageMetric?url=' + uri, true);//第二步：打开连接  将请求参数写在url中  ps:"./Ptest.php?name=test&nameone=testone"
metricRequest.send();//第三步：发送请求  将请求参数写在URL中


/**
 * 获取数据后的处理程序
 */
pageMetricHttpRequest.onreadystatechange = function () {
    if (pageMetricHttpRequest.readyState == 4 && pageMetricHttpRequest.status == 200) {
        let json = pageMetricHttpRequest.responseText;//获取到json字符串，还需解析
        console.log(json);

        let ps = document.getElementsByTagName('p');
        for(let i=0; i < ps.length; i++) {
            if (ps[i].hasAttribute("wk-blog-component-pageMetric") ) {
                ps[i].innerText = json;
            }

            if (ps[i].hasAttribute("wk-blog-component-siteMetric") ) {
                ps[i].innerText = json;
            }
        }

    }
};
