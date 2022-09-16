package main

import "fmt"

func main() {
	var t, n, m int
	fmt.Scan(&t)
	fmt.Scan(&n)
	v := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&v[i])
	}
	fmt.Scan(&m)
	w := make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Scan(&w[i])
	}

	dp := make([]int, t+10)
	for i := 1; i <= t; i++ {
		dp[i] = -1
	}
	for i := range w {
		for j := t; j >= w[i]; j-- {
			if dp[j-w[i]] == -1 {
				continue
			}
			dp[j] = max(dp[j], dp[j-w[i]]+v[i])
		}
	}

	ans := 0
	for i := range dp {
		ans = max(ans, dp[i])
	}

	fmt.Println(ans)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//group1/M00/00/00/CgPIJmMh2O6ALP-yAAAAQa56jaM8074419

/*

configure arguments: --with-cc-opt='-g -O2 -fstack-protector-strong -Wformat -Werror=format-security -fPIC -Wdate-time -D_FORTIFY_SOURCE=2' --with-ld-opt='-Wl,-z,relro -Wl,-z,now -fPIC' --prefix=/usr/share/nginx --conf-path=/etc/nginx/nginx.conf --http-log-path=/var/log/nginx/access.log --error-log-path=/var/log/nginx/error.log --lock-path=/var/lock/nginx.lock --pid-path=/run/nginx.pid --modules-path=/usr/lib/nginx/modules --http-client-body-temp-path=/var/lib/nginx/body --http-fastcgi-temp-path=/var/lib/nginx/fastcgi --http-proxy-temp-path=/var/lib/nginx/proxy --http-scgi-temp-path=/var/lib/nginx/scgi --http-uwsgi-temp-path=/var/lib/nginx/uwsgi --with-debug --with-pcre-jit --with-http_ssl_module --with-http_stub_status_module --with-http_realip_module --with-http_auth_request_module --with-http_v2_module --with-http_dav_module --with-http_slice_module --with-threads --with-http_addition_module --with-http_geoip_module=dynamic --with-http_gunzip_module --with-http_gzip_static_module --with-http_image_filter_module=dynamic --with-http_sub_module --with-http_xslt_module=dynamic --with-stream=dynamic --with-stream_ssl_module --with-stream_ssl_preread_module --with-mail=dynamic --with-mail_ssl_module  --add-module=/root/work/blockchain/fastdfs/fastdfs-nginx-module-1.22/src

*/
