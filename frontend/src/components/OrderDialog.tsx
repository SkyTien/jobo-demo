import { Dialog, DialogContent, DialogTitle, IconButton } from "@mui/material";
import CloseIcon from '@mui/icons-material/Close';
import PatientForm from "./PatientForm";

export default function OrderDialog(props: any) {
  const { onClose, selectedValue, open } = props;
  
    return (
      <Dialog open={open}>
        <DialogTitle sx={{ m: 0, p: 2 }}>
        {selectedValue?.name}
        {onClose ? (
          <IconButton
            aria-label="close"
            onClick={onClose}
            sx={{
              position: 'absolute',
              right: 8,
              top: 8,
              color: (theme) => theme.palette.grey[500],
            }}
          >
            <CloseIcon />
          </IconButton>
          ) : null}
      </DialogTitle>
      <DialogContent dividers>
        <PatientForm orderId={selectedValue?.orderId}/>
      </DialogContent>
    </Dialog>
    );
}