-- Seed data for electronics products
INSERT INTO products (id, name, description, price, stock, category, image_url, is_active, created_at, updated_at)
VALUES
  (gen_random_uuid(), 'iPhone 15 Pro', 'Apple flagship smartphone with A17 chip, 128GB storage, 6.1" display.', 1199.99, 20, 'Phones', '/uploads/iphone15pro.jpg', true, NOW(), NOW()),
  (gen_random_uuid(), 'Samsung Galaxy S24 Ultra', 'Samsung premium phone, 256GB, 200MP camera, 6.8" AMOLED.', 1099.99, 15, 'Phones', '/uploads/galaxys24ultra.jpg', true, NOW(), NOW()),
  (gen_random_uuid(), 'MacBook Pro 16" M3', 'Apple M3 Pro chip, 16GB RAM, 1TB SSD, 16" Liquid Retina.', 2499.99, 10, 'Laptops', '/uploads/macbookpro16m3.jpg', true, NOW(), NOW()),
  (gen_random_uuid(), 'Dell XPS 13', '13.4" FHD+, Intel i7, 16GB RAM, 512GB SSD, Windows 11.', 1399.99, 12, 'Laptops', '/uploads/dellxps13.jpg', true, NOW(), NOW()),
  (gen_random_uuid(), 'Sony Bravia XR 65" OLED', '65" 4K OLED, Google TV, XR Cognitive Processor, HDMI 2.1.', 1999.99, 8, 'TVs', '/uploads/sonybraviaxr65.jpg', true, NOW(), NOW()),
  (gen_random_uuid(), 'LG C3 55" OLED', '55" 4K OLED, webOS, Dolby Vision, HDMI 2.1, 120Hz.', 1499.99, 10, 'TVs', '/uploads/lgc355.jpg', true, NOW(), NOW()),
  (gen_random_uuid(), 'Apple iPad Air (2024)', '10.9" Liquid Retina, M2 chip, 128GB, Wi-Fi.', 699.99, 18, 'Tablets', '/uploads/ipadair2024.jpg', true, NOW(), NOW()),
  (gen_random_uuid(), 'Samsung Galaxy Tab S9', '11" AMOLED, Snapdragon 8 Gen 2, 256GB, S Pen.', 799.99, 14, 'Tablets', '/uploads/galaxytabs9.jpg', true, NOW(), NOW()),
  (gen_random_uuid(), 'Sony WH-1000XM5', 'Wireless Noise Cancelling Headphones, 30h battery.', 349.99, 25, 'Accessories', '/uploads/sonywh1000xm5.jpg', true, NOW(), NOW()),
  (gen_random_uuid(), 'Apple Watch Series 9', 'GPS, 45mm, Always-On Retina, Blood Oxygen, ECG.', 429.99, 16, 'Accessories', '/uploads/applewatchs9.jpg', true, NOW(), NOW());
