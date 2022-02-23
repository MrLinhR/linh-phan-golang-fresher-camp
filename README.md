# Update bằng cách
- Tạo thêm function trong storage để update mỗi khi insert hoặc delete 1 bản ghi.
- Tại tầng storage ta dùng thư viện GORM để update giá trị cho cột
- Tạo thêm 1 store cho tầng business để gọi hàm tăng, giảm giá trị cột đếm.
# Bằng cách
- Dùng middleware để tránh crash
