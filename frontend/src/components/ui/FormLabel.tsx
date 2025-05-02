import React from 'react';
import { Text, TextProps } from '@chakra-ui/react';

interface FormLabelProps extends TextProps {
  htmlFor?: string;
}

// A simple custom FormLabel component that doesn't rely on Chakra UI's FormLabel
const FormLabel: React.FC<FormLabelProps> = ({ children, htmlFor, ...props }) => {
  return (
    <Text
      as="label"
      htmlFor={htmlFor}
      fontWeight="medium"
      mb="2px"
      display="block"
      {...props}
    >
      {children}
    </Text>
  );
};

export default FormLabel;