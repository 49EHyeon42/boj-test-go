## 그래프 구현

- 인접 행렬 그래프
- 인접 정렬 그래프

## 그래프 방향

- 무방향 그래프
- 유방향 그래프

## 그래프 가중치

- 무가중치 그래프
- 유가중치 그래프

## 그래프 복잡도

- 공간 복잡도
  - 인접 행렬 : O(v^2)
  - 인접 리스트 : O(V+E)
- 시간 복잡도
  - 두 노드가 연결되었는지 확인하는데 걸리는 시간
    - 인접 행렬 : O(1)
    - 인접 리스트 : O(V)
  - 한 노드에 연결된 모든 노드들을 확인하는데 걸리는 시간
    - 인접 행렬 : O(V)
    - 인접 리스트 : O(E)
