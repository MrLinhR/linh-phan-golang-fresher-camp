## Index trong database : ##  
- Là một dạng cấu trúc dùng để xác định vị trí và truy cập vào dữ liệu trong các bảng database
- Tối ưu hiệu suất truy vấn database bằng việc giảm lượng truy cập vào bộ nhớ khi thực hiện truy vấn
- Index thường được tạo mặc định cho primary key, foreign key.
- Index là một cấu trúc dữ liệu gồm:
    + Cột Search Key: chứa bản sao các giá trị của cột được tạo Index
    + Cột Data Reference: chứa con trỏ trỏ đến địa chỉ của bản ghi có giá trị cột index tương ứng

## Khác biệt khi có và không có index: ##  
- Tiết kiệm thời gian khi tra cứu dữ liệu