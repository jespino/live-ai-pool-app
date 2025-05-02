import React from 'react';
import { Box, Button, Center, Heading, Text, VStack, Link } from '@chakra-ui/react';
import { QRCodeSVG } from 'qrcode.react';
import { usePool } from '../contexts/PoolContext';
import { ExternalLinkIcon } from '@chakra-ui/icons';

interface PoolQRCodeProps {
  onContinue: () => void;
}

const PoolQRCode: React.FC<PoolQRCodeProps> = ({ onContinue }) => {
  const { currentPool, loading, error } = usePool();
  
  // Generate the URL for voting
  const voteUrl = currentPool 
    ? `${window.location.origin}/vote/${currentPool.id}`
    : '';

  if (loading) {
    return (
      <Center h="100vh">
        <Text>Loading pool...</Text>
      </Center>
    );
  }

  if (error || !currentPool) {
    return (
      <Center h="100vh">
        <VStack spacing={4}>
          <Text color="red.500">{error || 'No active pool found'}</Text>
          <Button onClick={onContinue}>Continue</Button>
        </VStack>
      </Center>
    );
  }

  return (
    <Box maxW="md" mx="auto" mt={8} p={6} borderWidth={1} borderRadius="lg" boxShadow="md">
      <VStack spacing={6}>
        <Heading size="lg" textAlign="center">Scan to Vote</Heading>
        <Text textAlign="center">{currentPool.title}</Text>
        <Text fontSize="sm" color="gray.600" textAlign="center">
          {currentPool.description}
        </Text>
        
        <Center p={4} bg="white" borderRadius="md" boxShadow="sm">
          <QRCodeSVG
            value={voteUrl}
            size={250}
            includeMargin
            level="H"
          />
        </Center>
        
        <VStack spacing={1}>
          <Text fontSize="sm" color="gray.600" textAlign="center">
            Scan this QR code with your phone to vote
          </Text>
          
          <Link 
            href={voteUrl} 
            color="blue.500" 
            isExternal
            fontWeight="medium"
            fontSize="md"
            display="flex"
            alignItems="center"
          >
            {voteUrl} <ExternalLinkIcon mx="2px" />
          </Link>
        </VStack>
        
        <Button colorScheme="blue" onClick={onContinue}>
          Continue to Results
        </Button>
      </VStack>
    </Box>
  );
};

export default PoolQRCode;