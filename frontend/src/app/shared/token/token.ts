import { jwtDecode } from 'jwt-decode';

function GetToken(): string | null {
    const cookieString = document.cookie;
    const cookies = cookieString.split(';');

    for (let i = 0; i < cookies.length; i++) {
        const cookie = cookies[i].trim();
        if (cookie.startsWith('auth_token=')) {
            return cookie.substring('auth_token='.length, cookie.length);
        }
    }

    return null;
}
function GetUserRole(): string | null {
    const token = GetToken();
    if (!token) return null;
    try {
        const decoded: any = jwtDecode(token);
        return decoded.role || null;
    } catch (e) {
        return null;
    }
}
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
function IsLoggedIn(): boolean {
    return !!GetToken();
}
export { GetToken, GetUserRole, SaveTokenToCookie, IsLoggedIn };
