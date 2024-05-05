'use client';

import {PageTitle} from '@/features/PageTitle';
import Layout from '@/shared/layouts/Layout';
import {Pie} from 'react-chartjs-2';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';

ChartJS.register(ArcElement, Tooltip, Legend);

export default function StatisticsPage() {
    const data = {
        labels: ['Buy', 'Sell'],
        datasets: [
            {
                label: '# of Votes',
                data: [12, 19 ],
                backgroundColor: [
                    'rgb(224,33,73)',
                    'rgb(54,235,132)',
                ],
                borderColor: [
                    'rgba(255, 99, 132, 1)',
                    'rgba(54, 162, 235, 1)',
                ],
                borderWidth: 1,
            },
        ],
    };

    return (
        <Layout>
            <PageTitle title={'Статистика'}/>
            <div style={{width: '180px', height: '180px'}}>
                <Pie data={data}/>
            </div>
        </Layout>
    );
}