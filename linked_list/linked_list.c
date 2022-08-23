#include <stdio.h>
#include <stdlib.h>

/*
 * The data structure that each element of the linked list consists of.
 */
typedef struct linked_list_node {
  int i;
  struct linked_list_node *child;
} linked_list_node;

/*
 * The wrapper that makes up a whole linked list.
 */
typedef struct linked_list {
  linked_list_node *head;
  linked_list_node *tail;
} linked_list;

/*
 * Add an element to the end of a linked list.
 */
void linked_list_append(linked_list *l, int i) {
  linked_list_node *n = malloc(sizeof(linked_list_node));
  n->child = NULL;
  n->i = i;
  l->tail->child = n;
  l->tail = n;
}

/*
 * Print out the contents of a linked list.
 */
void linked_list_print(linked_list_node *n) {
  if (!n) {
    return;
  }
  printf("%d\n", n->i);
  linked_list_print(n->child);
}

/*
 * Free every element in a linked list.
 */
void free_all(linked_list_node *n) {
  if (!n) {
    return;
  }
  free_all(n->child);
  free(n);
}

/*
 * Initialize a linked list struct.
 */
linked_list *linked_list_init() {
  linked_list_node *n = malloc(sizeof(linked_list_node));
  n->child = NULL;
  n->i = 0;
  linked_list *l = malloc(sizeof(linked_list));
  l->head = n;
  l->tail = n;

  return l;
}

int main() {
  linked_list *l = linked_list_init();

  /*
   * Add ten elements to the linked list.
   */
  for (int i = 0; i < 10; i++) {
    linked_list_append(l, i);
  }

  /*
   * Print out the linked list.
   */
  linked_list_print(l->head->child);

  /*
   * Cleanup.
   */
  free_all(l->head);
  free(l);
}
