import { Box, Button, TextField } from "@mui/material";
import React, { useEffect, useReducer } from "react";
import { Order } from "./Model";

type State = {
    order: Order | {},
    isEdit: boolean
};

type Action = {
    type: ActionMode,
    payload?: any
}

enum ActionMode { 
    EDIT_ORDER,
    SET_ORDER,
    SAVE_ORDER
};

const initState: State = {
    order: {},
    isEdit: false
};

const reducer = (state: State, action: Action) => {
    switch(action.type) {
        case ActionMode.EDIT_ORDER:
            return {
                ...state,
                isEdit: !state.isEdit,
            };
        case ActionMode.SET_ORDER: 
            return {
                ...state,
                order: action.payload,
            }
        case ActionMode.SAVE_ORDER:
            return {
                ...state,
                isEdit: false
            };
      default:
        throw new Error(`Unhandled action type: ${action.type}`);
    }
}

export default function PatientForm(props: { orderId: number }) {
    const { orderId } = props;

    const [state, dispatch] = useReducer(reducer, initState);

    useEffect(() => {
        fetch('/api/patient/order/' + orderId)
           .then((response) => response.json())
           .then((data) => {
                dispatch({type: ActionMode.SET_ORDER, payload: data});
           })
           .catch((err) => {
              console.log(err.message);
           });
     }, []);

    const save = () => {
        fetch('/api/patient/order/' + orderId, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ message: state.order.message }) // convert data to JSON string
            })
            .then(response => response.json())
            .then(data => {
            console.log('Success:', data);
            })
            .catch((error) => {
            console.error('Error:', error);
            });
            dispatch({type: ActionMode.SAVE_ORDER});
    }

    return (
            <Box
                component="form"
                sx={{
                '& .MuiTextField-root': { m: 1, width: '100%' },
                }}
                noValidate
                autoComplete="off"
            >
            <TextField
                fullWidth
                id="outlined-read-only-input"
                multiline
                maxRows={4}
                label='Message'
                value={state.order.message || ''}
                onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                    dispatch({type: ActionMode.SET_ORDER, payload: {id: state.order.id, message: event.target.value}})
                }}
                InputProps={{
                    readOnly: !state.isEdit,
                    disabled: !state.isEdit
                }}
            />
            {state.isEdit && (
                <Button variant="outlined" onClick={save}>SAVE</Button>
            )}
            {
            !state.isEdit && (
                <Button variant="outlined" onClick={() => dispatch({type: ActionMode.EDIT_ORDER})}>EDIT</Button>
            )}
        </Box>
    );
}