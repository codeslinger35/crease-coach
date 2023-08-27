import { Typography, Grid, Divider, Box, List, ListItemButton, ListItemText, Link } from '@mui/material';
import * as React from 'react';

export default function GameDetails({game}) {
  const [selectedIndex, setSelectedIndex] = React.useState(0);
  const [coachingPoint, setCoachingPoint] = React.useState({});
  const handleListItemClick = (event, coachingPoint) => {
    setSelectedIndex(coachingPoint.id);
    setCoachingPoint(coachingPoint);
  };

  const savePct = (game) => {
    let pct = game.periods.reduce(((prev, curr) => prev += curr.saves ), 0)/game.periods.reduce(((prev, curr) => prev += curr.shotsAgainst ), 0);
    return pct.toPrecision(4);
  }

  return (
    <div>
      <Grid container spacing="1" rowSpacing='2'>
        <Grid item xs="4">
          <Typography className='centered-text' variant='body'>Opponent: {game.opponent}</Typography>
        </Grid>
        <Grid item xs="4">
          <Typography className='centered-text' variant='body'>Date: {game.date}</Typography>
        </Grid>
        <Grid item xs="4">
          <Typography className='centered-text' variant='body'>Shots againt: {game.periods.reduce(((prev, curr) => prev += curr.shotsAgainst ), 0)}</Typography>
          <br></br>
          <Typography className='centered-text' variant='body'>Save %: {savePct(game)}</Typography>
        </Grid>

        <Grid item xs='12'>
          <Divider />
        </Grid>

        {game.periods.map((p) => (
          <Grid item xs={12/game.periods.length}>
            <Typography variant='body'>Period #: {p.periodNumber}</Typography>
            <br></br>
            <Typography variant='body'>Notes: {p.notes}</Typography>
          </Grid>
        ))}

        <Grid item xs='12'></Grid>

        <Grid item xs='4'>
          <Box sx={{width: '100%', maxWidth: 360}}>
            <List component="nav" style={{maxHeight:256, overflow: 'auto'}}>
              {game.periods.map((p) => p.coachingPoints.map((cp) => (
                <ListItemButton
                key={cp.id}
                selected={selectedIndex === cp.id}
                onClick={(event) => handleListItemClick(event, cp)}
                >
                  <ListItemText primary={cp.time} secondary={cp.event}></ListItemText>
                </ListItemButton>
              ))) }
            </List>
          </Box>
        </Grid>

        <Grid item xs='8'>
          <Grid container spacing='1'>
          <Grid item xs='4'>
              <Typography className='centered-text' variant='body'>Event: {coachingPoint.event}</Typography>
            </Grid>
            <Grid item xs='4'>
              <Typography className='centered-text' variant='body'>Time: {coachingPoint.time}</Typography>
            </Grid>
            <Grid item xs='4'>
              <Link className='centered-text' variant='body' target="_blank" rel="noopener" href={coachingPoint.url}>Hudl</Link>
            </Grid>
            <Divider />
            <Grid item xs='12'>
              <Box>
                <Typography className='centered-text' variant='body'>Notes: {coachingPoint.notes}</Typography>
              </Box>
            </Grid>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
}
