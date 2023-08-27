import { Card, CardContent } from "@mui/material";
import { Typography } from "@mui/material";

function PlayerCard({firstName, lastName, age, team}) {
  return (
    <Card className="centered">
      <CardContent>
      <Typography variant="h2" component="h1">
        {firstName} {lastName}
      </Typography>
      <Typography variant="h3" component="h3">
        {team}
      </Typography>
      <Typography variant="h4" component="h4">
        {age}
      </Typography>
      </CardContent>
    </Card>
  )
}

export default PlayerCard;
