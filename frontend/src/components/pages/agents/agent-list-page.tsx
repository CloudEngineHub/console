import {
  Button,
  ButtonGroup,
  DataTable,
  Empty,
  Flex,
  HStack,
  Icon,
  SearchField,
  Spinner,
  Stack,
  Text,
} from '@redpanda-data/ui';
import { runInAction } from 'mobx';
import { Pipeline_State } from 'protogen/redpanda/api/dataplane/v1/pipeline_pb';
// import {
//   ListPipelinesRequest as ListPipelinesRequestDataPlane,
//   ListPipelinesRequest_Filter,
//   Pipeline_State,
// } from 'protogen/redpanda/api/dataplane/v1/pipeline_pb';

import { useEffect, useState } from 'react';
import { AiOutlineDelete } from 'react-icons/ai';
import { FaRegStopCircle } from 'react-icons/fa';
import { MdCheck, MdDone, MdError, MdOutlineQuestionMark, MdRefresh } from 'react-icons/md';
import { useListPipelinesQuery } from 'react-query/api/pipeline';
import { useHistory } from 'react-router-dom';
import { uiState } from 'state/uiState';

// Hack for MobX to ensure we don't need to use observables
export const updatePageTitle = () => {
  runInAction(() => {
    uiState.pageTitle = 'Agents';
    uiState.pageBreadcrumbs.pop(); // Remove last breadcrumb to ensure the agent title is used without previous page breadcrumb being shown
    uiState.pageBreadcrumbs.push({ title: 'Agents', linkTo: '/agents', heading: 'Agents' });
  });
};

interface AgentStateDisplayValueProps {
  status?: Pipeline_State;
}

const AgentStateDisplayValue = ({ status }: AgentStateDisplayValueProps) => {
  switch (status) {
    case Pipeline_State.UNSPECIFIED: {
      return (
        <HStack spacing={2}>
          <Icon as={MdOutlineQuestionMark} color="red" />
          <Text>Unknown</Text>
        </HStack>
      );
    }
    case Pipeline_State.STARTING: {
      return (
        <HStack spacing={2}>
          <Icon as={MdRefresh} color="blue" />
          <Text>Starting</Text>
        </HStack>
      );
    }
    case Pipeline_State.RUNNING: {
      return (
        <HStack spacing={2}>
          <Icon as={MdCheck} color="green" />
          <Text>Running</Text>
        </HStack>
      );
    }
    case Pipeline_State.STOPPING: {
      return (
        <HStack spacing={2}>
          <Icon as={MdRefresh} color="blue" />
          <Text>Stopping</Text>
        </HStack>
      );
    }
    case Pipeline_State.STOPPED: {
      return (
        <HStack spacing={2}>
          <Icon as={FaRegStopCircle} />
          <Text>Stopped</Text>
        </HStack>
      );
    }
    case Pipeline_State.ERROR: {
      return (
        <HStack spacing={2}>
          <Icon as={MdError} color="red" />
          <Text>Error</Text>
        </HStack>
      );
    }
    case Pipeline_State.COMPLETED: {
      return (
        <HStack spacing={2}>
          <Icon as={MdDone} color="green" />
          <Text>Completed</Text>
        </HStack>
      );
    }
  }
};

export const AgentListPage = () => {
  const history = useHistory();

  const [nameContains, setNameContains] = useState('');
  const {
    data: agentList,
    isLoading: isAgentListLoading,
    isError: isAgentListError,
  } = useListPipelinesQuery(
    // new ListPipelinesRequestDataPlane({
    //   filter: nameContains
    //     ? new ListPipelinesRequest_Filter({
    //         nameContains,
    //         tags: {
    //           type: 'agent',
    //         },
    //       })
    //     : undefined,
    // }),
  );

  console.log('agentList: ', agentList);

  // const filteredAgentList = agentList?.fil

  useEffect(() => {
    updatePageTitle();
  }, []);

  if (isAgentListError) {
    return <Empty />;
  }

  return (
    <Stack spacing={8}>
      <Stack spacing={4}>
        <Text>Manage your AI agents.</Text>
        <ButtonGroup>
          <Button
            variant="outline"
            onClick={() => {
              history.push('/agents/create');
            }}
            data-testid="create-agent-button"
          >
            Create new Agent
          </Button>
        </ButtonGroup>
      </Stack>

      <SearchField
        width="350px"
        searchText={nameContains}
        setSearchText={setNameContains}
        placeholderText="Filter agents..."
      />

      {isAgentListLoading ? (
        <Flex justifyContent="center" padding={8}>
          <Spinner size="lg" />
        </Flex>
      ) : agentList?.pipelines?.length === 0 ? (
        <Empty />
      ) : (
        <DataTable
          data={agentList?.pipelines ?? []}
          pagination
          defaultPageSize={10}
          sorting
          columns={[
            {
              header: 'Name',
              cell: ({ row: { original } }) => (
                <Text data-testid={`agent-name-${original?.id}`}>{original?.displayName}</Text>
              ),
              size: 200,
            },
            {
              header: 'Description',
              cell: ({ row: { original } }) => (
                <Text data-testid={`agent-description-${original?.id}`}>{original?.description}</Text>
              ),
              size: 200,
            },
            {
              header: 'Status',
              id: 'status',
              cell: ({ row: { original } }) => <AgentStateDisplayValue status={original?.state} />,
              size: 50,
            },
            {
              header: '',
              id: 'actions',
              cell: ({ row: { original } }) => (
                <HStack spacing={4} justifyContent="flex-end" width="100%">
                  <Icon
                    data-testid={`delete-agent-${original?.id}`}
                    as={AiOutlineDelete}
                    // onClick={() => handleDeleteSecretModal(original?.id ?? '')}
                    cursor="pointer"
                    aria-label="Delete agent"
                  />
                </HStack>
              ),
            },
          ]}
        />
      )}
    </Stack>
  );
};
