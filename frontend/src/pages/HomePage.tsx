import { Box, Button, Container, Flex, Heading } from '@chakra-ui/react';
import React, { useState } from 'react';
import PoolQRCode from '../components/PoolQRCode';
import PoolResults from '../components/PoolResults';

enum HomeView {
  QR_CODE,
  RESULTS
}

const HomePage: React.FC = () => {
  const [view, setView] = useState<HomeView>(HomeView.QR_CODE);

  return (
    <Container maxW="container.xl" py={8}>
      <Flex justifyContent="space-between" alignItems="center" mb={6}>
        <Heading>Pool App</Heading>
      </Flex>

      {view === HomeView.QR_CODE ? (
        <PoolQRCode onContinue={() => setView(HomeView.RESULTS)} />
      ) : (
        <Box>
          <Button 
            mb={4} 
            onClick={() => setView(HomeView.QR_CODE)}
            colorScheme="blue"
            variant="outline"
          >
            Back to QR Code
          </Button>
          <PoolResults />
        </Box>
      )}
    </Container>
  );
};

export default HomePage;