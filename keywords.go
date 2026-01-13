resp, err := grequests.Patch("http:/microsoft.com/test?NTID=1",
            &grequests.RequestOptions{
                IsAjax: true, 
            })

resp, err := grequests.Patch("http:/microsoft.com/test?account_identifier=1",
            &grequests.RequestOptions{
                IsAjax: true, 
            })

resp, err := grequests.Patch("http:/microsoft.com/test?bankruptcy=true",
            &grequests.RequestOptions{
                IsAjax: true, 
            })

resp, err := grequests.Patch("http:/microsoft.com/test?billing_address=123",
            &grequests.RequestOptions{
                IsAjax: true, 
            })
