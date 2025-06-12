import {
  Box,
  Typography,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow
} from '@mui/material';

// Моковые данные расходов
const monthlyCosts = [
  {
    month: 'Январь',
    maintenance: 120000,
    security: 50000,
    cleaning: 30000,
    utilities: 90000,
    extras: 20000
  },
  {
    month: 'Февраль',
    maintenance: 100000,
    security: 52000,
    cleaning: 28000,
    utilities: 85000,
    extras: 0
  },
  {
    month: 'Март',
    maintenance: 110000,
    security: 51000,
    cleaning: 29000,
    utilities: 88000,
    extras: 15000
  }
];

export function OperatingCostsAnalytics() {
  return (
    <Box sx={{ mt: 3 }}>
      <Typography variant="h5" gutterBottom>
        Операционные расходы по месяцам
      </Typography>

      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Месяц</TableCell>
              <TableCell align="right">Обслуживание</TableCell>
              <TableCell align="right">Охрана</TableCell>
              <TableCell align="right">Уборка</TableCell>
              <TableCell align="right">Коммунальные</TableCell>
              <TableCell align="right">Доп. расходы</TableCell>
              <TableCell align="right">Итого</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {monthlyCosts.map((cost, i) => {
              const total =
                cost.maintenance +
                cost.security +
                cost.cleaning +
                cost.utilities +
                cost.extras;
              return (
                <TableRow key={i}>
                  <TableCell>{cost.month}</TableCell>
                  <TableCell align="right">{cost.maintenance.toLocaleString()}</TableCell>
                  <TableCell align="right">{cost.security.toLocaleString()}</TableCell>
                  <TableCell align="right">{cost.cleaning.toLocaleString()}</TableCell>
                  <TableCell align="right">{cost.utilities.toLocaleString()}</TableCell>
                  <TableCell align="right">{cost.extras.toLocaleString()}</TableCell>
                  <TableCell align="right">{total.toLocaleString()}</TableCell>
                </TableRow>
              );
            })}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
}
