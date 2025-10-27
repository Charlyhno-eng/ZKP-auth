"use client";

import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import { useState } from "react";
import RegisterModal from "./RegisterModal";

export default function Navbar() {
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);

  return (
    <AppBar
      sx={{
        backgroundColor: "transparent",
        borderBottom: "1px solid rgba(255, 255, 255, 0.12)",
        boxShadow: 0,
      }}
    >
      <Toolbar
        sx={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
        }}
      >
        <Typography
          variant="h5"
          component="div"
          sx={{ color: "white", fontWeight: 700 }}
        >
          navbar
        </Typography>
        <Box>
          <Button color="inherit">Se connecter</Button>
          <Button color="inherit" onClick={handleOpen}>
            Créer un compte
          </Button>
          <RegisterModal open={open} onClose={handleClose} />
        </Box>
      </Toolbar>
    </AppBar>
  );
}
