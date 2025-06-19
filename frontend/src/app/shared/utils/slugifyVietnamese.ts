export function SlugifyVietnamese(text: string): string {
    return text
        .toLowerCase()
        .normalize("NFD") // tách dấu ra khỏi ký tự gốc
        .replace(/[\u0300-\u036f]/g, "") // xoá các dấu
        .replace(/đ/g, "d") // thay thế đ
        .replace(/[^a-z0-9\s-]/g, "") // xoá ký tự đặc biệt
        .trim() // xoá khoảng trắng đầu/cuối
        .replace(/\s+/g, "-"); // thay khoảng trắng bằng gạch ngang
}

