import axios from "axios";
const instance = axios.create({
    // 这边记得修改你对应的配置文件
    baseURL:  "http://localhost:8080",
    withCredentials: true
})


instance.interceptors.response.use(function (resp) {
    const newToken = resp.headers["x-jwt-token"]
    console.log("resp headers", resp.headers)
    console.log("token" + newToken)
    if (newToken) {
        localStorage.setItem("token", newToken)
    }
    if (resp.status == 401) {
        window.location.href="/users/login"
    }
    return resp
}, (err) => {
    console.log(err)
    if (err.response.status == 401) {
        window.location.href="/users/login"
    }
    return err
})

// 在这里让每一个请求都加上 authorization 的头部
instance.interceptors.request.use((req) => {
    const token = localStorage.getItem("token")
    req.headers.setAuthorization("Bearer " + token, true)
    return req
}, (err) => {
    console.log(err)
})

export default instance