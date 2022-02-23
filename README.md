# Khi nào cần tạo các cột số đếm ngay trên table dữ liệu (VD: liked_count trên restaurants)?
- Khi cần đếm số lượng của dữ liệu nào đó ở db
- Ở trường hợp liked_count số lượt like liên quan đến việc người dùng có thể unlike, việc đếm ở db sẽ làm cho db phải chịu tải, nên ta tạo 1 bảng riêng, module riêng rồi cout ở đó  