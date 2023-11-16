document.getElementById('loginForm').addEventListener("submit", function(event){
    event.preventDefault()//阻止表单的默认提交

    //console.log(event.target)
    const formData = new FormData(event.target)//获取表单数据
    //console.log(formData)

    sendLoginRequest(formData)//发送POST请求
})

function sendLoginRequest(formData){
    const xhr = new XMLHttpRequest()

    xhr.responseType = 'json';
    xhr.open("POST", "/manage/manage-login", true)

    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4 && xhr.status === 200){
            const resp = xhr.response
            if (xhr.response.code === 200){
                //console.log(response)
                const token = resp.data.token

                //存储token
                //console.log(token)
                localStorage.setItem("token", token)

                alert("登录成功！")

                //登录成功后的逻辑，例如重定向受保护的页面
                window.location.href = "/manage"
            }
            else {
                alert("登录失败!")
                console.log("登录失败，错误信息：", xhr.response.message)
            }
        } else if (xhr.readyState === 4){
            //处理登录失败的情况
            alert("登录失败，请检查用户名和密码")
        }
    }
    //发送请求
    xhr.send(formData)
}