import { Box, Typography } from '@mui/material';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer, Brush } from 'recharts';

const mockMonthlyIncome = [
  { month: 'Янв', income: 120000 },
  { month: 'Фев', income: 135000 },
  { month: 'Мар', income: 150000 },
  { month: 'Апр', income: 142000 },
  { month: 'Май', income: 160000 },
];

export function IncomeAnalytics() {
  return (
    <Box p={2}>
      <Typography variant="h6" gutterBottom>Доходы по месяцам</Typography>

      <ResponsiveContainer width="100%" height={300}>
        <LineChart data={mockMonthlyIncome}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="month" />
          <YAxis />
          <Tooltip />
          <Line type="monotone" dataKey="income" stroke="#1976d2" strokeWidth={2} />
          <Brush dataKey="month" height={30} stroke="#1976d2" />
        </LineChart>
      </ResponsiveContainer>
    </Box>
  );
}
