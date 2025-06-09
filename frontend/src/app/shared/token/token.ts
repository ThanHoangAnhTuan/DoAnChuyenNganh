function SaveTokenToCookie(token: string) {
    // Tham số của document.cookie: name=value; expires=date; path=path; domain=domain; secure

    // Thiết lập thời gian hết hạn (1h)
    const expirationDate = new Date();
    expirationDate.setTime(expirationDate.getTime() + 1 * 60 * 60 * 1000);

    // Thiết lập cookie với các tùy chọn bảo mật
    document.cookie = `auth_token=${token}; expires=${expirationDate.toUTCString()}; path=/; SameSite=Strict`;

    // Nếu sử dụng HTTPS, bạn có thể thêm thuộc tính 'secure'
    // document.cookie = `auth_token=${token}; expires=${expirationDate.toUTCString()}; path=/; SameSite=Strict; secure`;
}

export {
    SaveTokenToCookie
}