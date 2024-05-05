'use client';

import {PageTitle} from '@/features/PageTitle';
import Layout from '@/shared/layouts/Layout';
import {Pie, PieChart, ResponsiveContainer} from 'recharts';

export default function StatisticsPage() {
    return (
        <Layout>
            <PageTitle title={'Статистика'}/>
            <div>
                <ResponsiveContainer width="100%" height={400}>
                    <PieChart width={400} height={400}>
                    <Pie dataKey={'value'} fill="#8884d8" innerRadius={60} outerRadius={80}
                         data={[{value: 10, name: 'A'}, {value: 15, name: 'B'}, {value: 20, name: 'C'}]}/>
                </PieChart>
                </ResponsiveContainer>
            </div>
        </Layout>
    );
}