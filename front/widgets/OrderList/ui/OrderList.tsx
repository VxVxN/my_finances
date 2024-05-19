'use client'

import React, {FC} from 'react';
import {mockOrders, Order, OrderCard} from '@/entities/OrderCard';
import { motion } from "framer-motion";
import './styles.scss'

interface OrderListProps {
}

const item = {
    hidden: { opacity: 0, x: "100%" },
    visible: {
      opacity: 1, 
      x: "0%"
    }
  };

export const OrderList: FC<OrderListProps> = ({}) => {
    const orders: Order[] = mockOrders

    return (
        <ul className="order-list">
            {orders.map((order: Order) => (
                <motion.li key={order.id} variants={item}  initial="hidden"
                animate="visible">
                  <OrderCard data={order}/> 
                </motion.li>
            ))}
        </ul>
    );
};