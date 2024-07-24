import React from 'react';
import { useMemo } from 'react';
import {MaterialReactTable} from 'material-react-table';
import { QueryClient, QueryClientProvider, useQuery } from '@tanstack/react-query';
import { Box, CircularProgress } from '@mui/material';
import { useNavigate } from 'react-router-dom';
function useGetInventory() {
  const inventoryid = localStorage.getItem('inventorysheetid')
  const ownerid = localStorage.getItem('uuid')
  return useQuery({
    queryKey: ['inventoryValues', inventoryid, ownerid],
      queryFn: async () => {
      const response = await fetch(`http://br-avus.com:8000/sheetapi/getValues?id=${inventoryid}&uuid=${ownerid}`);
      if (!response.ok) {
          throw new Error('Network response was not ok');
      }
      return response.json();
      },
      refetchOnWindowFocus: false,
  });
}

const InventorySheet = () => {
  const { data, isError, isLoading } = useGetInventory();
  const navigate = useNavigate();
  // Transform the data
  const columns = useMemo(() => {
      if (!data?.valueRanges?.[0]?.values) return [];
  const headers = data.valueRanges[0].values[0];
  return headers
    .map((header, index) => (header.trim() ? { accessorKey: `column${index}`, header: header.trim() } : null))
    .filter(Boolean);
}, [data]);
  
    const tableData = useMemo(() => {
      if (!data?.valueRanges?.[0]?.values) return [];
      return data.valueRanges[0].values.slice(1).map((row) => {
        const rowData = {};
        row.forEach((cell, index) => {
          rowData[`column${index}`] = cell;
        });
        return rowData;
      });
    }, [data]);
if (data == null) {
  navigate('/importsheet')
}
if (isLoading) {
  return (
    <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
      <CircularProgress />
    </Box>
  );
}

if (isError) {
  return <div>Error loading data</div>;
}

return (
  <Box sx={{ overflowX: 'auto' }}>
    <MaterialReactTable
      columns={columns}
      data={tableData}
      initialState={{ columnVisibility: {} }}
      enableStickyHeader
      enableColumnResizing
    />
  </Box>
);
};

export default InventorySheet;