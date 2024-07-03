import React from 'react';
import { useMemo } from 'react';
import {MaterialReactTable} from 'material-react-table';
import { QueryClient, QueryClientProvider, useQuery } from '@tanstack/react-query';
import { Box, CircularProgress } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import InventorySheet from './InventorySheet';
import FinancialSheet from './FinancialSheet';


const queryClient = new QueryClient();

const MainSheet = () => (
  <QueryClientProvider client={queryClient}>1
  <div>
    <FinancialSheet/>
  </div>
  <dvi>
    <InventorySheet />
  </dvi>
  </QueryClientProvider>
);

export default MainSheet;