import { useEffect, useState } from 'react';
import { Button, Grid } from "@mui/material";
import { useNavigate, useParams } from "react-router-dom";
import PlayerCard from './PlayerCard';
import SeasonList from './SeasonList';
import Games from './Games';

function GoalieDetail() {
  const navigate = useNavigate();
  const [goalie, setData] = useState({});
  const [games, setGames] = useState([]);
  const params = useParams();

  const getData = (id) => {
    fetch('../goalies.json', { headers: {'Content-Type': "application/json", 'Accept': "application/json"}})
    .then(rawData => { return rawData.json()})
    .then(list => { return list.find(g => g.id === parseInt(id))})
    .then(json => setData(json));
  }

  const getGames = (season) => {
    setGames(season.games);
  }

  useEffect(() => {
    getData(params.id);
  }, [params]);
  return (
    <div className="table-container">
      <Button
        onClick={() => navigate(-1) }
      >
        Back
      </Button>
      <PlayerCard className="centered"
          firstName={goalie.firstName}
          lastName={goalie.lastName}
          team={goalie.team}
          age={goalie.age}
        />
        <Grid container direction="row" spacing={1}>
          <Grid item xs={3}>
            <SeasonList seasons={goalie.seasons ?? []} seasonSelect={getGames}/>
          </Grid>
          <Grid item xs={9} style={{minHeight: 256}}>
            <Games games={games} />
          </Grid>
        </Grid>
    </div>
  )
}

export default GoalieDetail;
