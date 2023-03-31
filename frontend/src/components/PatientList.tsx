import { Box, List, ListItem, ListItemButton, ListItemText } from "@mui/material";
import React, { useCallback, useEffect, useState } from "react";
import { Patient } from "./Model";
import OrderDialog from "./OrderDialog";

export default function PatientList() {
    const [open, setOpen] = React.useState(false);
    const [patients, setPatients] = useState([]);
    const [selectedValue, setSelectedValue] = React.useState<Patient | undefined>(undefined);
    const handleClickOpen = useCallback(
        (item: Patient) => () => {
            setOpen(true);
            setSelectedValue(item)
        },
        [],
      );
  
    const handleClose = (value: string) => {
      setOpen(false);
    };

    useEffect(() => {
        fetch('/api/patient/lists')
           .then((response) => response.json())
           .then((data) => {
              setPatients(data);
           })
           .catch((err) => {
              console.log(err.message);
           });
     }, []);    
    
    return (
        <Box sx={{
            width: 300,
            height: 300,
            border: '1px dashed grey',
            borderRadius: 5
          }}>
            <List style={{width: '100%'}}>
                {
                patients.map((item: Patient) =>
                <ListItem key={item.id} disablePadding>
                    <ListItemButton onClick={handleClickOpen(item)}>
                        <ListItemText primary={item.name} />
                    </ListItemButton>
                </ListItem>)
            }
            </List>
            <OrderDialog selectedValue={selectedValue} open={open} onClose={handleClose}/>
      </Box>
    );
}