import {Order, OrderType} from '@/entities/OrderCard';

export const mockOrders: Order[] =
    [
        {
            "id": 1,
            "type": OrderType.BUY,
            "datetime": new Date("2024-05-05T08:30:18.295953552Z"),
            baseCurrency: "WETH",
            "value": 95,
            quoteCurrency: "USD",
            "price": 3135
        },
        {
            "id": 2,
            "type": OrderType.BUY,
            "datetime": new Date("2024-05-05T08:30:18.295953552Z"),
            baseCurrency: "TON",
            "value": 52,
            quoteCurrency: "USD",
            "price": 3449
        },
        {
            "id": 3,
            "type": OrderType.BUY,
            "datetime": new Date("2024-05-05T08:30:18.295953552Z"),
            baseCurrency: "ETH",
            "value": 94,
            quoteCurrency: "USD",
            "price": 3105
        },
        {
            "id": 4,
            "type": OrderType.BUY,
            "datetime": new Date("2024-05-05T08:30:18.295953552Z"),
            baseCurrency: "WETH",
            "value": 47,
            quoteCurrency: "USD",
            "price": 3451
        },
        {
            "id": 5,
            "type": OrderType.BUY,
            "datetime": new Date("2024-05-05T08:30:18.295953552Z"),
            baseCurrency: "WETH",
            "value": 83,
            quoteCurrency: "USD",
            "price": 3360
        },
        {
            "id": 6,
            "type": OrderType.SELL,
            "datetime": new Date("2024-05-05T08:30:18.295953552Z"),
            baseCurrency: "TON",
            "value": 63,
            quoteCurrency: "USD",
            "price": 3245
        },
        {
            "id": 7,
            "type": OrderType.BUY,
            "datetime": new Date("2024-05-05T08:30:18.295953552Z"),
            baseCurrency: "WETH",
            "value": 82,
            quoteCurrency: "USD",
            "price": 3312
        },
        {
            "id": 8,
            "type": OrderType.BUY,
            "datetime": new Date("2024-05-05T08:30:18.295953552Z"),
            baseCurrency: "NANO",
            "value": 70,
            quoteCurrency: "USD",
            "price": 3189
        },
        {
            "id": 9,
            "type": OrderType.BUY,
            "datetime": new Date("2024-05-05T08:30:18.295953552Z"),
            baseCurrency: "WETH",
            "value": 45,
            quoteCurrency: "USD",
            "price": 3141
        },
        {
            "id": 10,
            "type": OrderType.BUY,
            "datetime": new Date("2024-05-05T08:30:18.295953552Z"),
            baseCurrency: "XER",
            "value": 36,
            quoteCurrency: "USD",
            "price": 3249
        }
    ]
