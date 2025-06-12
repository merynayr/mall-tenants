import { Box, Typography, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from '@mui/material';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';

// Моковые данные для графика
const debtTrend = [
  { month: 'Янв', totalDebt: 120000 },
  { month: 'Фев', totalDebt: 140000 },
  { month: 'Мар', totalDebt: 180000 },
  { month: 'Апр', totalDebt: 250000 },
  { month: 'Май', totalDebt: 300000 },
  { month: 'Июн', totalDebt: 420000 },
];

// Моковые данные для таблицы
const debtors = [
  { tenant: 'ООО Альфа', amount: 150000, daysOverdue: 45, unit: '1-101' },
  { tenant: 'ИП Бета', amount: 120000, daysOverdue: 30, unit: '2-202' },
  { tenant: 'ООО Гамма', amount: 100000, daysOverdue: 15, unit: '3-305' },
  { tenant: 'ЗАО Дельта', amount: 50000, daysOverdue: 90, unit: '1-110' },
];

export function DebtsAnalytics() {
  return (
    <Box sx={{ mt: 3 }}>
      <Typography variant="h5" gutterBottom>
        Общая задолженность: 420 000 ₽
      </Typography>

      <Typography variant="subtitle1" gutterBottom>
        Динамика задолженности по месяцам
      </Typography>
      <Box height={300}>
        <ResponsiveContainer width="100%" height="100%">
          <LineChart data={debtTrend}>
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey="month" />
            <YAxis />
            <Tooltip />
            <Line type="monotone" dataKey="totalDebt" stroke="#f44336" strokeWidth={2} />
          </LineChart>
        </ResponsiveContainer>
      </Box>

      <Typography variant="subtitle1" sx={{ mt: 4, mb: 1 }}>
        Текущие должники
      </Typography>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Арендатор</TableCell>
              <TableCell>Помещение</TableCell>
              <TableCell>Сумма задолженности</TableCell>
              <TableCell>Дней просрочки</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {debtors.map((row, i) => (
              <TableRow key={i}>
                <TableCell>{row.tenant}</TableCell>
                <TableCell>{row.unit}</TableCell>
                <TableCell>{row.amount.toLocaleString()} ₽</TableCell>
                <TableCell>{row.daysOverdue}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
}
