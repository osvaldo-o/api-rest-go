CREATE TABLE orders (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,           -- ID único, asumimos que es numérico.
    status BOOLEAN NOT NULL,                        -- Representa el estado como verdadero/falso.
    name VARCHAR(255) NOT NULL,                     -- Nombre del cliente.
    phone VARCHAR(20) NOT NULL,                     -- Número de teléfono.
    delivery_date DATE NOT NULL,                    -- Fecha de entrega.
    delivery_time TIME NOT NULL,                    -- Hora de entrega.
    place_delivery VARCHAR(255) NOT NULL,           -- Lugar de entrega.
    description TEXT,                               -- Descripción del pedido (puede ser NULL).
    price DECIMAL(10, 2) NOT NULL,                  -- Precio, formato numérico con 2 decimales.
    comment TEXT                                    -- Comentarios adicionales (puede ser NULL).
);

INSERT INTO orders (status, name, phone, delivery_date, delivery_time, place_delivery, description, price, comment)
VALUES
(true, 'Alice Johnson', '5551234567', '2025-02-01', '09:00:00', '742 Evergreen Terrace', 'First order with a special request.', 150.75, 'Handle with care'),
(false, 'Bob Smith', '5559876543', '2025-02-05', '15:30:00', '221B Baker Street', 'Order for testing delivery delay.', 89.50, 'Delivery time is flexible'),
(true, 'Charlie Brown', '5556547890', '2025-02-10', '12:15:00', '742 Evergreen Terrace', 'Repeat customer order.', 120.00, 'Leave at the front door'),
(true, 'Diana Prince', '5554561230', '2025-02-15', '17:45:00', 'Themyscira Island', 'Bulk order for a family gathering.', 999.99, 'Call before delivery'),
(false, 'Ethan Hunt', '5557891234', '2025-03-01', '08:30:00', 'Mission Impossible HQ', 'Urgent delivery.', 200.00, 'Delivery should be top secret');