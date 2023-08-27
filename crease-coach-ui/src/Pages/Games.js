import * as React from 'react';
import { DataGrid } from "@mui/x-data-grid";
import { Box, Modal } from '@mui/material';
import GameDetails from './GameDetails';

const columns = [
  { field: 'date', headerName: "Date" },
  { field: 'opponent', headerName: "Opponent" },
  { headerName: "Save %", valueGetter: (params) => {
    let shotsAgainst = params.row.periods.reduce(((prev, current) => prev += current.shotsAgainst), 0);
    let saves = params.row.periods.reduce(((prev, current) => prev += current.saves), 0);
    return (saves/shotsAgainst).toPrecision(4);
  }},
];

const style = {
  position: 'absolute',
  top: '10%',
  left: '10%',
  bottom: '10%',
  right: '10%',
  transform: 'translate(-10%, -10%, -10%, -10%)',
  bgcolor: 'background.paper',
  border: '2px solid #000',
  boxShadow: 24,
  p: 4,
};

function Games({games}) {

  const viewGame = (
    params, // GridRowParams
    event, // MuiEvent<React.MouseEvent<HTMLElement>>
    details, // GridCallbackDetails
  ) => {
    setGame(params.row);
    setOpen(true);
  };

  const [game, setGame] = React.useState({});
  const [open, setOpen] = React.useState(false);
  const handleClose = () => setOpen(false);

  return (
    <div className="table-container">
      <DataGrid
        columns={columns}
        rows={games}
        onRowClick={viewGame}
      />
      <div>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <GameDetails game={game}/>
        </Box>
      </Modal>
    </div>
    </div>
  );
}

export default Games;
