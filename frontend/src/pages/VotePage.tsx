import React, { useEffect } from 'react';
import { Box, Container, Heading, Text, Button, VStack, Center } from '@chakra-ui/react';
import { useParams, Link } from 'react-router-dom';
import VotingForm from '../components/VotingForm';
import { usePool } from '../contexts/PoolContext';
import { getPool } from '../services/api';

const VotePage: React.FC = () => {
  const { poolId } = useParams<{ poolId: string }>();
  const { currentPool, loading, error, refreshPool } = usePool();

  // Load the correct pool when the page loads
  useEffect(() => {
    const loadPoolData = async () => {
      if (poolId && (!currentPool || currentPool.id !== poolId)) {
        try {
          await refreshPool();
        } catch (err) {
          console.error('Error loading pool:', err);
        }
      }
    };

    loadPoolData();
  }, [poolId, currentPool, refreshPool]);

  // Check if the current pool matches the requested poolId
  const isCorrectPool = currentPool && currentPool.id === poolId;

  if (loading) {
    return (
      <Container maxW="container.md" py={8}>
        <Center>
          <Text>Loading pool...</Text>
        </Center>
      </Container>
    );
  }

  if (error || !isCorrectPool) {
    return (
      <Container maxW="container.md" py={8}>
        <VStack spacing={6}>
          <Heading>Invalid Pool</Heading>
          <Text color="red.500">
            {error || 'The requested poll does not exist or is no longer active.'}
          </Text>
          <Link to="/">
            <Button colorScheme="blue">Return to Home</Button>
          </Link>
        </VStack>
      </Container>
    );
  }

  return (
    <Container maxW="container.md" py={8}>
      <Box mb={6}>
        <Heading textAlign="center" mb={2}>Vote Now</Heading>
      </Box>

      <VotingForm />

      <Box mt={6} textAlign="center">
        <Link to="/">
          <Button variant="outline">Return to Home</Button>
        </Link>
      </Box>
    </Container>
  );
};

export default VotePage;