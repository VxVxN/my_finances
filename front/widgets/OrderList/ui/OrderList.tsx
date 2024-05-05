import React, {FC} from 'react';
import {mockOrders, Order, OrderCard} from '@/entities/OrderCard';
import './styles.scss'

interface OrderListProps {
    // orders: Order[];
}

export const OrderList: FC<OrderListProps> = ({}) => {
    const orders: Order[] = mockOrders

    return (
        <ul className="order-list">
            {orders.map((order: Order) => (
                <li key={order.id}>
                  <OrderCard data={order}/>
                </li>
            ))}
        </ul>
    );
};