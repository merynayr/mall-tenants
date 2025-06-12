import { Box, Typography, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, LinearProgress } from '@mui/material';

// Моковые данные по этажам
const occupancyByFloor = [
  { floor: '1 этаж', total: 1000, rented: 800 },
  { floor: '2 этаж', total: 1200, rented: 900 },
  { floor: '3 этаж', total: 900, rented: 300 },
];

export function OccupancyAnalytics() {
  const totalArea = occupancyByFloor.reduce((sum, f) => sum + f.total, 0);
  const totalRented = occupancyByFloor.reduce((sum, f) => sum + f.rented, 0);
  const totalOccupancyPercent = Math.round((totalRented / totalArea) * 100);

  return (
    <Box sx={{ mt: 3 }}>
      <Typography variant="h5" gutterBottom>
        Общая заполненность: {totalOccupancyPercent}%
      </Typography>

      <LinearProgress
        variant="determinate"
        value={totalOccupancyPercent}
        sx={{ height: 10, borderRadius: 5, mb: 4 }}
      />

      <Typography variant="subtitle1" sx={{ mb: 1 }}>
        Заполненность по этажам
      </Typography>

      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Этаж</TableCell>
              <TableCell align="right">Общая площадь, м²</TableCell>
              <TableCell align="right">Сдано в аренду, м²</TableCell>
              <TableCell align="right">Заполненность</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {occupancyByFloor.map((floor, i) => {
              const percent = Math.round((floor.rented / floor.total) * 100);
              return (
                <TableRow key={i}>
                  <TableCell>{floor.floor}</TableCell>
                  <TableCell align="right">{floor.total}</TableCell>
                  <TableCell align="right">{floor.rented}</TableCell>
                  <TableCell align="right">{percent}%</TableCell>
                </TableRow>
              );
            })}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
}
