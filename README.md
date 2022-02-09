## Không nên chứa file upload vào ngay chính bên trong service mà nên dùng Cloud bởi: ##
- Hệ thống cần có tính linh hoạt, lưu trữ ảnh tại hệ thống khiến hệ thống nặng, tiêu tốn tài nguyên.  
- Có tính bảo mật cao, dễ quản lý.
- Giúp ảnh hoạt động trên nhiều hệ thống.
  
## Không chứa binary ảnh vào DB bởi: ##
- Gây nặng DB, lãng phí tài nguyên và gây mất thời gian khi truy xuất
