import { useState } from "react";
import { Box, List, ListItemButton, ListItemText, Typography } from "@mui/material";

function SeasonList({seasons, seasonSelect}) {

  const [selectedIndex, setSelectedIndex] = useState(0);

  const handleListItemClick = (event, season, index) => {
    setSelectedIndex(index);
    seasonSelect(season);
  };

  return (
    <Box sx={{width: '100%', maxWidth: 360}}>
      <Typography variant='h5' style={{padding: 8}}>Seasons</Typography>
      <List component="nav">
        {
          seasons.map((s, index) => (
            <ListItemButton
              key={s.id}
              selected={selectedIndex === index}
              onClick={(event) => handleListItemClick(event, s, index)}
            >
              <ListItemText primary={s.title}/>
          </ListItemButton>
          ))
        }
      </List>
    </Box>
  )
}

export default SeasonList;
