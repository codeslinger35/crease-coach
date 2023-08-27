import { useEffect, useState } from 'react';
import { DataGrid } from '@mui/x-data-grid';
import { useNavigate } from 'react-router-dom';
import { Typography } from '@mui/material';

const tableColumns = [
  { field: 'firstName', headerName: 'First name' },
  { field: 'lastName', headerName: 'Last name' },
  { field: 'team', headerName: 'Current Team' },
  { field: 'age', headername: 'Age group'}
];

function Search() {
  const getData = () => {
    fetch('goalies.json', { headers: {'Content-Type': "application/json", 'Accept': "application/json"}})
    .then(rawData => { return rawData.json()})
    .then(json => setData(json));
  }

  const viewDetails = (
    params, // GridRowParams
    event, // MuiEvent<React.MouseEvent<HTMLElement>>
    details, // GridCallbackDetails
  ) => {
    navigate('goalie/'+params.row.id)
  };

  useEffect(() => {
    getData();
  });

  const [data, setData] = useState([]);
  const navigate = useNavigate();

  return (
    <div className='table-container'>
      <Typography variant="h1" component="h1">
        Crease Coach - Carolina Jr Hurricanes
      </Typography>
      <DataGrid
        columns={tableColumns}
        rows={data}
        onRowClick={viewDetails}
      />
    </div>
  );
}

export default Search;
