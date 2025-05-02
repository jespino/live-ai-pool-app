import React, { useState } from 'react';
import { 
  Box, 
  Button, 
  Heading, 
  Radio, 
  RadioGroup, 
  Stack, 
  Text, 
  VStack, 
  useToast 
} from '@chakra-ui/react';
import { usePool } from '../contexts/PoolContext';

const VotingForm: React.FC = () => {
  const { currentPool, loading, error, voted, vote } = usePool();
  const [selectedOption, setSelectedOption] = useState<string>('');
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
  const toast = useToast();

  if (loading) {
    return (
      <Box maxW="md" mx="auto" mt={8} p={6} borderWidth={1} borderRadius="lg" boxShadow="md">
        <Text textAlign="center">Loading pool...</Text>
      </Box>
    );
  }

  if (error || !currentPool) {
    return (
      <Box maxW="md" mx="auto" mt={8} p={6} borderWidth={1} borderRadius="lg" boxShadow="md">
        <Text color="red.500" textAlign="center">
          {error || 'No active pool found'}
        </Text>
      </Box>
    );
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!selectedOption) return;
    
    setIsSubmitting(true);
    
    try {
      await vote(selectedOption);
      toast({
        title: 'Vote submitted',
        description: 'Your vote has been recorded',
        status: 'success',
        duration: 3000,
        isClosable: true,
      });
    } catch (err) {
      console.error('Voting failed:', err);
    } finally {
      setIsSubmitting(false);
    }
  };

  if (voted) {
    return (
      <Box maxW="md" mx="auto" mt={8} p={6} borderWidth={1} borderRadius="lg" boxShadow="md">
        <VStack spacing={4}>
          <Heading size="md">Thank You!</Heading>
          <Text>Your vote has been recorded.</Text>
          <Text fontWeight="bold">Pool: {currentPool.title}</Text>
        </VStack>
      </Box>
    );
  }

  return (
    <Box maxW="md" mx="auto" mt={8} p={6} borderWidth={1} borderRadius="lg" boxShadow="md">
      <form onSubmit={handleSubmit}>
        <VStack spacing={6}>
          <Heading size="md" textAlign="center">{currentPool.title}</Heading>
          <Text textAlign="center">{currentPool.description}</Text>

          <RadioGroup onChange={setSelectedOption} value={selectedOption} width="100%">
            <Stack direction="column" spacing={4}>
              {currentPool.options.map((option) => (
                <Radio key={option.id} value={option.id}>
                  {option.text}
                </Radio>
              ))}
            </Stack>
          </RadioGroup>

          <Button
            type="submit"
            colorScheme="blue"
            width="full"
            isLoading={isSubmitting}
            isDisabled={!selectedOption}
          >
            Submit Vote
          </Button>
        </VStack>
      </form>
    </Box>
  );
};

export default VotingForm;