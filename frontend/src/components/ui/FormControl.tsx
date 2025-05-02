import React from 'react';
import { Box, BoxProps } from '@chakra-ui/react';

interface FormControlProps extends BoxProps {
  isInvalid?: boolean;
}

// A simple custom FormControl component that doesn't rely on Chakra UI's FormControl
const FormControl: React.FC<FormControlProps> = ({ children, isInvalid, ...props }) => {
  return (
    <Box
      role="group"
      mb={4}
      {...props}
      data-invalid={isInvalid ? 'true' : undefined}
    >
      {children}
    </Box>
  );
};

export default FormControl;