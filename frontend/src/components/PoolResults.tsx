import React from 'react';
import { 
  Box, 
  Button, 
  Flex, 
  Heading, 
  HStack, 
  Progress, 
  Text, 
  VStack, 
  Badge
} from '@chakra-ui/react';
import { usePool } from '../contexts/PoolContext';

const PoolResults: React.FC = () => {
  const { currentPool, loading, error, nextPool, previousPool } = usePool();

  if (loading) {
    return (
      <Box maxW="lg" mx="auto" mt={8} p={6} borderWidth={1} borderRadius="lg" boxShadow="md">
        <Text textAlign="center">Loading results...</Text>
      </Box>
    );
  }

  if (error || !currentPool) {
    return (
      <Box maxW="lg" mx="auto" mt={8} p={6} borderWidth={1} borderRadius="lg" boxShadow="md">
        <Text color="red.500" textAlign="center">
          {error || 'No active pool found'}
        </Text>
      </Box>
    );
  }

  // Calculate total votes and percentages
  const totalVotes = currentPool.options.reduce((sum, option) => sum + option.votes, 0);
  
  // Sort options by votes (descending)
  const sortedOptions = [...currentPool.options].sort((a, b) => b.votes - a.votes);

  // Format dates
  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleString();
  };

  return (
    <Box maxW="lg" mx="auto" mt={8} p={6} borderWidth={1} borderRadius="lg" boxShadow="md">
      <VStack spacing={6} align="stretch">
        <Heading size="lg" textAlign="center">{currentPool.title}</Heading>
        <Text textAlign="center">{currentPool.description}</Text>
        
        <HStack justifyContent="center" spacing={4}>
          <Badge colorScheme="blue">Total Votes: {totalVotes}</Badge>
          <Badge colorScheme="green">Created: {formatDate(currentPool.created_at)}</Badge>
          <Badge colorScheme="orange">Expires: {formatDate(currentPool.expires_at)}</Badge>
        </HStack>

        <VStack spacing={4} align="stretch">
          {sortedOptions.map((option) => {
            const percentage = totalVotes === 0 ? 0 : Math.round((option.votes / totalVotes) * 100);
            
            return (
              <Box key={option.id} p={3} borderWidth={1} borderRadius="md">
                <Text fontWeight="bold" mb={1}>
                  {option.text}
                </Text>
                <Flex align="center">
                  <Progress 
                    value={percentage} 
                    size="md" 
                    colorScheme="blue" 
                    borderRadius="md"
                    flex="1" 
                    mr={3}
                  />
                  <Text fontWeight="bold" minWidth="80px" textAlign="right">
                    {option.votes} ({percentage}%)
                  </Text>
                </Flex>
              </Box>
            );
          })}
        </VStack>

        <HStack justifyContent="center" spacing={4} mt={4}>
          <Button onClick={previousPool} colorScheme="gray">
            Previous Pool
          </Button>
          <Button onClick={nextPool} colorScheme="blue">
            Next Pool
          </Button>
        </HStack>
      </VStack>
    </Box>
  );
};

export default PoolResults;