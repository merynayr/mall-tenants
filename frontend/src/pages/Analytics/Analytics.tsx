import { Tabs, Tab, Box } from '@mui/material';
import { IncomeAnalytics } from '@/components/Analytics/IncomeAnalytics';
import { DebtsAnalytics } from '@/components/Analytics/DebtsAnalytics';
import { useState } from 'react';
import { OccupancyAnalytics } from '@/components/Analytics/OccupancyAnalytics';
import { OperatingCostsAnalytics } from '@/components/Analytics/OperatingCostsAnalytics';

export function PageAnalytics() {
  const [value, setValue] = useState(0);

  return (
    <Box sx={{ width: '100%' }}>
      <Tabs value={value} onChange={(_, newValue) => setValue(newValue)}>
        <Tab label="Доходы" />
        <Tab label="Задолженности" />
        <Tab label="Заполненность" />
        <Tab label="Расходы" />
      </Tabs>

      {value === 0 && <IncomeAnalytics />}
      {value === 1 && <DebtsAnalytics />}
      {value === 2 && <OccupancyAnalytics />}
      {value === 3 && <OperatingCostsAnalytics />}
    </Box>
  );
}