import { useState } from "react";
import { Box, List, ListItemButton, ListItemText } from "@mui/material";

function SeasonList({seasons, seasonSelect}) {

  const [selectedIndex, setSelectedIndex] = useState(0);

  const handleListItemClick = (event, season) => {
    setSelectedIndex(season.id);
    seasonSelect(season);
  };

  return (
    <Box sx={{width: '100%', maxWidth: 360}}>
      <List component="nav">
        {
          seasons.map((s) => (
            <ListItemButton
              key={s.id}
              selected={selectedIndex === s.id}
              onClick={(event) => handleListItemClick(event, s)}
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
