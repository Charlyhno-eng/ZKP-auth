import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Modal from "@mui/material/Modal";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import { useState } from "react";

type RegisterModalProps = {
  open: boolean;
  onClose: () => void;
};

const style = {
  position: "absolute" as const,
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  width: 400,
  bgcolor: "#242526",
  color: "#fff",
  border: "none",
  borderRadius: 2,
  boxShadow: 24,
  p: 4,
};

export default function RegisterModal({ open, onClose }: RegisterModalProps) {
  const [username, setUsername] = useState("");
  const [error, setError] = useState("");

  const handleRegister = async () => {
    setError("");
    const res = await fetch("http://localhost:8080/api/register", {
      method: "POST",
      body: JSON.stringify({ username }),
      headers: { "Content-Type": "application/json" },
    });
    if (res.ok) {
      const { private_key } = await res.json();
      const blob = new Blob([private_key], {
        type: "application/octet-stream",
      });
      const link = document.createElement("a");

      link.href = window.URL.createObjectURL(blob);
      link.download = "private_key.txt";
      link.click();

      onClose();
    } else {
      setError("Erreur lors de la création");
    }
  };

  return (
    <Modal open={open} onClose={onClose}>
      <Box sx={style}>
        <Typography variant="h6" component="h2" sx={{ mb: 2 }}>
          Créer un compte
        </Typography>
        <TextField
          label="Nom d'utilisateur"
          variant="filled"
          fullWidth
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          InputProps={{ sx: { background: "#35363a", color: "#fff" } }}
          InputLabelProps={{ sx: { color: "#bbb" } }}
          sx={{ mb: 2 }}
        />
        <Button
          variant="contained"
          fullWidth
          color="primary"
          sx={{ background: "#1976d2", color: "#fff" }}
          onClick={handleRegister}
        >
          Créer
        </Button>
        {error && (
          <Typography color="error" sx={{ mt: 2 }}>
            {error}
          </Typography>
        )}
      </Box>
    </Modal>
  );
}
