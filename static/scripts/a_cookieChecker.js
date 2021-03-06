function getCookie(name) {
    var cookie = " " + document.cookie;
    var search = " " + name + "=";
    var setStr = null;
    var offset = 0;
    var end = 0;
    if (cookie.length > 0) {
        offset = cookie.indexOf(search);
        if (offset !== -1) {
            offset += search.length;
            end = cookie.indexOf(";", offset);
            if (end === -1) {
                end = cookie.length;
            }
            setStr = unescape(cookie.substring(offset, end));
        }
    }
    return (setStr);
}

$(function() {
    var isLogin = getCookie("lib-login");
    if (isLogin != null && isLogin.length > 0) {
        window.isLogin = true;
    }
    var userId = getCookie("lib-id");
    if (userId != null && userId.length > 0) {
        window.userId = userId;
    }
});