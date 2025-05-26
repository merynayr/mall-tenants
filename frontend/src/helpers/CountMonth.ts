export function calculateRentalDuration(start: string, end: string): number | null {
  const startDate = new Date(start);
  const endDate = new Date(end);

  if (isNaN(startDate.getTime()) || isNaN(endDate.getTime()) || endDate <= startDate) {
    return null;
  }

  const startYear = startDate.getFullYear();
  const startMonth = startDate.getMonth();
  const startDay = startDate.getDate();

  const endYear = endDate.getFullYear();
  const endMonth = endDate.getMonth();
  const endDay = endDate.getDate();

  let months = (endYear - startYear) * 12 + (endMonth - startMonth);

  if (endDay > startDay) {
    months += 1;
  }

  return months === 0 ? 1 : months;
}