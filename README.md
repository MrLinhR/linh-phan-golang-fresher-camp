# Để giao tiếp với nhau:
- khi có nhiều module, muốn các module liên kết ta store dữ liệu của 1 module trong tầng storage của module này sau đó gọi và lấy dữ liệu rồi thao tác trong business của 1 module khác 

# Truy xuất:
- Store dữ liệu ở tầng storage ở đây là dùng để tính số lần nhà hàng được like, sau đó business của Restaurant gọi lại store này lấy dữ liệu
