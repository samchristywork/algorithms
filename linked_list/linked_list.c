#include <stdio.h>
#include <stdlib.h>

typedef struct linked_list_node {
  int i;
  struct linked_list_node *child;
} linked_list_node;

typedef struct linked_list {
  linked_list_node *head;
  linked_list_node *tail;
} linked_list;

void linked_list_append(linked_list *l, int i) {
  linked_list_node *n = malloc(sizeof(linked_list_node));
  n->child = NULL;
  n->i = i;
  l->tail->child = n;
  l->tail = n;
}

void linked_list_print(linked_list_node *n) {
  if (!n) {
    return;
  }
  printf("%d\n", n->i);
  linked_list_print(n->child);
}

void free_all(linked_list_node *n) {
  if (!n) {
    return;
  }
  free_all(n->child);
  free(n);
}

linked_list *linked_list_init(){
  linked_list_node *n = malloc(sizeof(linked_list_node));
  n->child = NULL;
  n->i = 0;
  linked_list *l = malloc(sizeof(linked_list));
  l->head = n;
  l->tail = n;

  return l;
}

int main() {
  linked_list *l=linked_list_init();

  for (int i = 0; i < 10; i++) {
    linked_list_append(l, i);
  }

  linked_list_print(l->head->child);

  free_all(l->head);
}
